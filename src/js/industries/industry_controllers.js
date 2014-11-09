var industryControllers = angular.module('industryControllers', []);

function IndustryListCtrl($scope, $http) {
  $http.get('/api/industries').success(function(data) {
    $scope.industries = data.results;
  });
  $scope.orderProp = 'name';
}

// Use the $inject property to prevent minification and munging from
// breaking the dependency injection
IndustryListCtrl.$inject = ['$scope', '$http'];
industryControllers.controller('IndustryListCtrl', IndustryListCtrl);

function IndustryDetailCtrl($scope, $routeParams, $http) {
  $http.get('/api/industries/' + $routeParams.pk).success(function(data) {
    $scope.industry = data;
  });
}

IndustryDetailCtrl.$inject = ['$scope', '$routeParams', '$http'];
industryControllers.controller('IndustryDetailCtrl', IndustryDetailCtrl);
