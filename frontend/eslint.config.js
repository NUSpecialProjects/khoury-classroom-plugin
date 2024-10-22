import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginReact from "eslint-plugin-react";

export default [
  { files: ["**/*.{js,jsx,mjs,cjs,ts,tsx}"] },
  { languageOptions: { globals: globals.browser } },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  pluginReact.configs.flat.recommended,
  {
    rules: { "react/react-in-jsx-scope": "off", "react/prop-types": "off" },
    settings: {
      react: {
        version: "detect",
      },
    },
  },
];
