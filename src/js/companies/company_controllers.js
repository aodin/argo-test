var companyControllers = angular.module('companyControllers', []);

function CompanyListCtrl($scope, $http) {
  $http.get('/api/companies').success(function(data) {
    $scope.companies = data.results;
  });
  $scope.orderProp = 'name';
}

// Use the $inject property to prevent minification and munging from
// breaking the dependency injection
CompanyListCtrl.$inject = ['$scope', '$http'];
companyControllers.controller('CompanyListCtrl', CompanyListCtrl);

function CompanyDetailCtrl($scope, $routeParams, $http) {
  $http.get('/api/companies/' + $routeParams.pk).success(function(data) {
    console.log(data);
    $scope.company = data;
  });
}

// Use the $inject property to prevent minification and munging from
// breaking the dependency injection
CompanyDetailCtrl.$inject = ['$scope', '$routeParams', '$http'];
companyControllers.controller('CompanyDetailCtrl', CompanyDetailCtrl);
