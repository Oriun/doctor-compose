/** @type {import('ts-jest/dist/types').InitialOptionsTsJest} */
module.exports = {
  preset: 'ts-jest',
  moduleFileExtensions: ['js','ts'],
  collectCoverage: true,
  coverageDirectory: './tests/coverage',
  testPathIgnorePatterns: ['node_modules', 'src', 'setup.test.ts'],
  setupFilesAfterEnv: [
    './tests/setup.test.ts'
  ],
  injectGlobals: true
};