module.exports = {
    extends: ['@commitlint/config-conventional'],
    rules: {
      'type-enum': [1, 'always', ['feat', 'fix', 'docs', 'test', 'wip']],
      'subject-case': [1, 'always', 'sentence-case'],
    },
  };
