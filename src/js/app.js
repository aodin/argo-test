var argoApp = angular.module('argoApp', [
  'ngRoute',
  'industryControllers'
]);

argoApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.when('/industries/:pk', {controller: 'IndustryDetailCtrl'})
  }
]);
