import Inquirer from "inquirer";
import { blue } from "./colors.js";
import fetch from "node-fetch";
import database_data from "./database_data.js";
import { randomBytes } from "crypto";

const SUPPORTED_DB = database_data.supported_db;
type DB = typeof SUPPORTED_DB[number];

const IMAGE_NAME = database_data.db_images as {
    [key in DB]: string
  };
const DB_STORAGE = database_data.db_storage as {
    [key in DB]: string
  };

const PERSIST_MODE = database_data.persistModes;
type PERSIST = typeof PERSIST_MODE[number];

const DOCKERHUB_TAGS = database_data.docker_hub_tags_url as {
  [key in DB]: string
};

const DEFAULT_PORT = database_data.default_ports as { [key: DB]: number };

type Env = {
  var_name: string;
  label: string;
  default: string;
  mandatory: boolean;
  description: string;
};
const DB_ENVS = database_data.db_envs as { [key in DB]: Env[] };

type TagsList = { latest100: string[]; recommended: string };
type TagsResponse = { results: { name: string }[] };
async function getTags(url: string): Promise<TagsList> {
  const response = await fetch(url);
  const body = (await response.json()) as TagsResponse;
  const allTags = body.results
    .map((tag: { name: string }) => tag.name)
    .sort()
    .reverse();
  const recommended =
    allTags.find((tag: string) => /^\d(\.\d(\.\d)?)$/.test(tag)) || allTags[0];
  return {
    latest100: allTags,
    recommended
  };
}

function populate(str: string, data?: { [key: string]: string }) {
  return str
  .replace(/\${RANDOM_STRING}/g, () => randomBytes(10).toString("base64").replace(/\=/g, ""))
  .replace(/\${([^}]+)}/g, (_, key) => data?.[key] || "");
}

export default async function Database() {
  console.log(blue("Let's create a fresh a database !"));

  var typeOfDatabase: DB,
    persistMode: PERSIST,
    persistLocation: string,
    persistVolume: string,
    expose: boolean,
    exposePort: number,
    restartPolicy: string,
    specificImageTag: string,
    loadEnvFromFile: boolean,
    env: { [key: string]: string | null } | null = {},
    serviceName: string,
    containerName: string;

  var { typeOfDatabase } = await Inquirer.prompt<{ typeOfDatabase: DB }>({
    type: "list",
    name: "typeOfDatabase",
    message: "What type of database are you creating ?",
    choices: SUPPORTED_DB
  });

  const images = getTags(DOCKERHUB_TAGS[typeOfDatabase]);

  var { persistMode } = await Inquirer.prompt<{ persistMode: PERSIST }>({
    type: "list",
    name: "persistMode",
    message: "Do you want to persist the database on disk or on a volume?",
    choices: PERSIST_MODE
  });

  if (persistMode === "Disk") {
    var { persistLocation } = await Inquirer.prompt<{
      persistLocation: string;
    }>({
      type: "input",
      name: "persistLocation",
      message: "Where to persist ?",
      default: "./" + typeOfDatabase.toLowerCase() + ".db",
      filter(input: string) {
        if(/^[a-z:\.]{0,3}\/.+/i.test(input)) return input
        return "./" + input
      }
    });
  } else {
    var { persistVolume } = await Inquirer.prompt<{ persistVolume: string }>({
      type: "input",
      name: "persistVolume",
      message: "Where to persist ?",
      default: typeOfDatabase.toLowerCase() + ".db"
    });
  }

  var { expose } = await Inquirer.prompt<{ expose: boolean }>({
    type: "confirm",
    name: "expose",
    message: "Do you want to expose it from outside of the app"
  });

  if (expose) {
    var { exposePort } = await Inquirer.prompt<{ exposePort: number }>({
      type: "number",
      name: "exposePort",
      message: "What port do you want to expose it on ?",
      default: DEFAULT_PORT[typeOfDatabase]
    });
  }

  var { restartPolicy } = await Inquirer.prompt<{ restartPolicy: string }>({
    type: "list",
    name: "restartPolicy",
    message: "Select the restart policy",
    choices: ["always", "on-failure", "unless-stopped", "never"],
    default: "unless-stopped"
  });

  const { latest100, recommended } = await images;
  console.log(blue("\nLoading image tags..."));
  var { specificImageTag } = await Inquirer.prompt<{
    specificImageTag: string;
  }>({
    type: "list",
    name: "specificImageTag",
    message: `Use a specific image tag ? (default to latest non-lts : ${recommended})`,
    choices: [recommended, ...latest100.filter(tag => tag !== recommended)],
    default: recommended,
    pageSize: 8
  });

  var { loadEnvFromFile } = await Inquirer.prompt<{
    loadEnvFromFile: boolean;
  }>({
    type: "confirm",
    name: "loadEnvFromFile",
    message: "Load enviroment variables from a file ?",
    default: true
  });

  for (const env_object of DB_ENVS[typeOfDatabase]) {
    if (loadEnvFromFile && !env_object.mandatory) {
      env[env_object.var_name] = null;
    } else {
      // TODO: add option to view env description
      const { env_value } = await Inquirer.prompt<{ env_value: string }>({
        type: "input",
        name: "env_value",
        message: populate(env_object.label),
        default: populate(env_object.default)
      });
      env[env_object.var_name] = env_value || null;
    }
  }

  var { serviceName } = await Inquirer.prompt<{ serviceName: string }>({
    type: "input",
    name: "serviceName",
    message: "What is the name of the service ?",
    default: "doctor-"+typeOfDatabase.toLowerCase()
  });

  var { containerName } = await Inquirer.prompt<{ containerName: string }>({
    type: "input",
    name: "containerName",
    message: "What is the name of the container ?",
    default: "doctor-"+typeOfDatabase.toLowerCase()
  });

  const service = {
    name: serviceName,
    "#description": typeOfDatabase + " database service",
    containerName,
    image: `${IMAGE_NAME[typeOfDatabase]}:${specificImageTag}`,
    volumes: [`${persistLocation! || persistVolume!}:${DB_STORAGE[typeOfDatabase]}`],
    restart: restartPolicy 
  } as { name: string, [key:string]: any}
  if(expose){
    service.ports = [`${exposePort!}:${DEFAULT_PORT[typeOfDatabase]}`]
  }
  if(!loadEnvFromFile){
    service.environment = env
    env = null
  }else {
    service.env_file = [".env"]
  }

  return {
    service,
    env
  }
}
