var argoApp = angular.module('argoApp', []);

function IndustryListCtrl($scope, $http) {
  $http.get('/api/industries').success(function(data) {
    $scope.industries = data.results;
  });
  $scope.orderProp = 'name';
}

// Use the $inject property to prevent minification and munging from
// breaking the dependency injection
IndustryListCtrl.$inject = ['$scope', '$http'];
argoApp.controller('IndustryListCtrl', IndustryListCtrl);
