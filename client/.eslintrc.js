module.exports = {
  root: true,
  env: {
    browser: true,
    node: true,
    es6: true
  },
  extends: ['plugin:vue/essential', '@vue/standard', '@vue/typescript'],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'space-before-function-paren': [2, 'never'],
    'vue/array-bracket-spacing': 'error',
    'vue/arrow-spacing': 'error',
    'vue/block-spacing': 'error',
    'vue/brace-style': 'error',
    'vue/camelcase': 'error',
    'vue/comma-dangle': 'error',
    'vue/component-name-in-template-casing': 'error',
    'vue/eqeqeq': 'error',
    'vue/key-spacing': 'error',
    'vue/match-component-file-name': 'error',
    'vue/object-curly-spacing': 'error'
  },
  parserOptions: {
    parser: '@typescript-eslint/parser',
    sourceType: 'module'
  }
}

// "eslintConfig": {
//   "root": true,
//   "env": {
//     "node": true
//   },
//   "extends": [
//     "plugin:vue/essential",
//     "@vue/prettier",
//     "@vue/typescript"
//   ],
//   "rules": {},
//   "parserOptions": {
//     "parser": "@typescript-eslint/parser"
//   }
// },
