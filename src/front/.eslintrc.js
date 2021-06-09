module.exports = {
  extends: [
    'plugin:vue/recommended',
    'airbnb-base'
  ],
  rules: {
    "vue/max-attributes-per-line": [
      "error",
      {
        "singleline": 5,
        "multiline": {
          "max": 1,
          "allowFirstLine": false
        }
      }
    ],
  },
}