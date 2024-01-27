module.exports = {
    extends: ['@commitlint/config-conventional'],
    rules: {
      'type-enum': [2, 'always', ['feat', 'fix', 'docs', 'test']],
      'subject-case': [2, 'always', 'sentence-case'],
    },
  };
