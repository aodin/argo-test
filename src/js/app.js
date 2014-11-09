var argoApp = angular.module('argoApp', [
  'ngRoute',
  'companyControllers',
  'industryControllers',
  'companyServices'
]);

argoApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/companies', {
        templateUrl: '/static/partials/company-list.html',
        controller: 'CompanyListCtrl'
      }).
      when('/companies/:pk', {
        templateUrl: '/static/partials/company-detail.html',
        controller: 'CompanyDetailCtrl'
      });
  }
]);
