import { writeFile } from "fs/promises";
import Inquirer from "inquirer";
import { blue } from "./src/colors.js";
import Database from "./src/database.js";
import Yml from "js-yaml";
import esMain from 'es-main';


const modules = {
  Database
} as {
  [key: string]: () => Promise<{
    service: { name: string; [key: string]: any };
    env: { [key: string]: string } | null;
  }>;
};

export async function writeCompose() {
  const services = {} as { [key: string]: any };
  const envs = [];
  const { typeOfApp } = (await Inquirer.prompt([
    {
      type: "list",
      name: "typeOfApp",
      message: "What type of app are you creating ?",
      choices: [...Object.keys(modules), "Other"]
    }
  ])) as { typeOfApp: keyof typeof modules | "Other" };

  if (typeOfApp === "Other") {
    console.log("Not Supported yet");
  } else {
    const { service: { name, ...conf }, env } = await modules[typeOfApp]();
    envs.push(env);
    services[name] = conf;
  }
  return {
    compose: {
      version: "3.9",
      services
    },
    env: {} as { [key: string]: string } | null
  };
}
export default async function main({} = {}) {
  console.log(
    blue(
      "\nWelcome to Doctor-Compose, the CLI that diagnose your app and find you the best docker-compose solution.\n"
    )
  );

  const { compose, env } = await writeCompose();

  await writeFile("docker-compose.yml", Yml.dump(compose).replace(/'#/g,'#'));
  if (env) {
    await writeFile(
      ".env",
      Object.entries(env)
        .map(([key, value]) => `${value ? "" : "#"}${key}=${value}`)
        .join("\n")
    );
  }
}
//@ts-ignore
if (esMain(import.meta)) {
    main()
} 
