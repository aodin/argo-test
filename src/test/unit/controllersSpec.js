var industriesResponse = {
  "meta": {
    "limit": 0,
    "offset": 0
  },
  "results": [
    {
      "about": "They're the future",
      "id": 1,
      "name": "Plastics"
    },
    {
      "about": "KAWAII!!!!",
      "id": 2,
      "name": "ハローキティ"
    },
    {
      "about": "Do you have them?",
      "id": 3,
      "name": "Battletoads"
    },
    {
      "about": "Lame",
      "id": 4,
      "name": "Venture Capital"
    }
  ]
}

describe('IndustryListCtrl', function(){
  var scope, ctrl, $httpBackend;

  beforeEach(module('argoApp'));

  beforeEach(inject(function(_$httpBackend_, $rootScope, $controller) {
    $httpBackend = _$httpBackend_;
    $httpBackend.expectGET('/api/industries').
        respond(industriesResponse);

    scope = $rootScope.$new();
    ctrl = $controller('IndustryListCtrl', {$scope: scope});
  }));

  it('should create "industries" model with ', function() {
    expect(scope.industries).toBeUndefined();
    $httpBackend.flush();

    expect(scope.industries.length).toEqual(4);
  });

  it('should set the default value of orderProp model', function() {
    expect(scope.orderProp).toBe('name');
  });

});