module.exports = function(config){
  config.set({

    basePath: '../../',

    files: [
      'bower_components/angular/angular.js',
      'bower_components/angular-route/angular-route.js',
      'bower_components/angular-resource/angular-resource.js',
      'bower_components/angular-mocks/angular-mocks.js',
      'src/js/**/*.js',
    ],

    autoWatch: true,
    frameworks: ['jasmine'],
    browsers: ['Firefox'],
    plugins : [
      'karma-chrome-launcher',
      'karma-firefox-launcher',
      'karma-jasmine'
    ],

    junitReporter : {
      outputFile: '_test/unit.xml',
      suite: 'unit'
    }

  });
};