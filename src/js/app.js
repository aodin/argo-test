var argoApp = angular.module('argoApp', [
  'ngRoute',
  'companyControllers',
  'industryControllers'
]);

argoApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.when('/industries/:pk', {controller: 'IndustryDetailCtrl'})
  }
]);
