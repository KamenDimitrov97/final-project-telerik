module.exports = {
    extends: ['@commitlint/config-conventional'],
    rules: {
      'type-enum': [2, 'always', ['feat', 'fix', 'docs', 'test', 'wip']],
      'subject-case': [2, 'always', 'sentence-case'],
    },
  };
