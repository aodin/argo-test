var companyServices = angular.module('companyServices', ['ngResource']);

companyServices.factory('Company', ['$resource',
  function($resource) {
    return $resource('/api/companies/:pk', {}, {
      'query': {method:'GET', isArray:false},
    });
  }
]);