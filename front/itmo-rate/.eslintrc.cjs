module.exports = {
  "env": {
    "browser": true,
    "es2023": true,
    "node": true
  },
  // "parser": "@typescript-eslint/parser",
  // "plugins":[
  //   "@typescript-eslint"
  // ],
  "extends": [
    "eslint:recommended",
    // "plugin:@typescript-eslint/eslint-recommended",
    // "plugin:@typescript-eslint/recommended",
    'plugin:vue/vue3-strongly-recommended',
    '@vue/eslint-config-typescript'
  ],
  "rules": {
    "vue/max-attributes-per-line": "off"
  },
  "parserOptions": {
    "ecmaVersion": "latest",
    "sourceType": "module"
  },
}