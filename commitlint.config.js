module.exports = {
    extends: ['@commitlint/config-conventional'],
    rules: {
      'type-enum': [48, 'always', ['feat', 'fix', 'docs', 'test', 'wip']],
      'subject-case': [48, 'always', 'sentence-case'],
    },
  };
