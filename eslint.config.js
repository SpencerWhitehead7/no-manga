// @ts-check

import eslint from "@eslint/js";
import prettierRecommended from "eslint-plugin-prettier/recommended";
import simpleImportSort from "eslint-plugin-simple-import-sort";
import tseslint from "typescript-eslint";

export default tseslint.config(
  {
    name: "universalRules",
    files: ["**/*.{js,ts}"],
    extends: [
      eslint.configs.recommended,
      tseslint.configs.strictTypeChecked,
      tseslint.configs.stylisticTypeChecked,
    ],
    plugins: {
      "simple-import-sort": simpleImportSort,
    },
    linterOptions: {
      reportUnusedDisableDirectives: "error",
    },
    rules: {
      "simple-import-sort/imports": "error",
      "simple-import-sort/exports": "error",
      "@typescript-eslint/restrict-template-expressions": [
        "error",
        { allowNumber: true },
      ],
      "@typescript-eslint/consistent-type-definitions": ["error", "type"],
    },
    languageOptions: {
      parserOptions: {
        project: "./tsconfig.json",
        // @ts-expect-error it so is defined tho
        tsconfigRootDir: import.meta.dirname,
      },
    },
  },
  {
    name: "disableTypeRulesInJs",
    files: ["**/*.js"],
    extends: [tseslint.configs.disableTypeChecked],
  },
  prettierRecommended,
);
