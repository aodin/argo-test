var companyControllers = angular.module('companyControllers', []);

function CompanyListCtrl($scope, Company) {
  $scope.companies = Company.query({}, function(companies) {});
  $scope.orderProp = 'name';
}

// Use the $inject property to prevent minification and munging from
// breaking the dependency injection
CompanyListCtrl.$inject = ['$scope', 'Company'];
companyControllers.controller('CompanyListCtrl', CompanyListCtrl);

function CompanyDetailCtrl($scope, $routeParams, Company) {
  $scope.company = Company.get({pk: $routeParams.pk}, function(phone) {});
}

// Use the $inject property to prevent minification and munging from
// breaking the dependency injection
CompanyDetailCtrl.$inject = ['$scope', '$routeParams', 'Company'];
companyControllers.controller('CompanyDetailCtrl', CompanyDetailCtrl);
