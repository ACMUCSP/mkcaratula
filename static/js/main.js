var caratulaApp = angular.module('CaratulaApp', [
    'ngRoute',
    'CaratulaControllers'
]);

caratulaApp.config(['$routeProvider',
    function ($routeProvider) {
        $routeProvider.
            when('/', {
                templateUrl: 'partials/main.html',
                controller: 'MainController'
            }).
            when('/tex', {
                templateUrl: 'partials/tex.html',
                controller: 'CodeController'
            }).
            when('/error', {
                templateUrl: 'partials/error.html',
                controller: 'CodeController'
            }).
            otherwise({
                redirectTo: '/'
            });
    }
]);

caratulaApp.filter('noscheme', function () {
    return function (input) {
        input = input || '';
        return input.slice(input.indexOf('://') + 3);
    };
});

var caratulaControllers = angular.module('CaratulaControllers', []);

caratulaControllers.controller('MainController', function ($scope, $http, $location, $rootScope) {
    $scope.loading = false;
    $scope.url = null;
    $scope.data = $rootScope.data ? $rootScope.data : {
        career: '',
        title: '',
        course: '',
        gender: 'm',
        names: [],
        semester: '1er'
    };
    if ($scope.data.names.length > 0) {
        $scope.names = [];
        for (var i = 0; i < $scope.data.names.length; ++i) {
            $scope.names.push({id: i + 1, name: $scope.data.names[i]})
        }
        $scope.data.names = [];
    } else {
        $scope.names = [{id: 1, name: ''}];
    }
    $scope.showNamesLabel = function (item) {
        return item.id === 1;
    };
    $scope.addNewNameBox = function (item) {
        if (item.name !== '' && item.id === $scope.names.length
            && $scope.names.length < 6) {
            $scope.names.push({id: $scope.names.length + 1, name: ''});
        }
    };
    $scope.removeNameBox = function (item) {
        if ($scope.names.length === 1) {
            item.name = '';
        } else {
            var cnt = 0, newNames = [];
            for (var i = 0; i < $scope.names.length; ++i) {
                var it = $scope.names[i];
                if (it !== item) {
                    it.id = ++cnt;
                    newNames.push(it);
                }
            }
            $scope.names = newNames;
        }
    };
    $scope.run = function (getLatexCode) {
        var url = "generate", newWindow;
        if ($scope.form.$invalid) {
            return;
        }
        for (var i = 0; i < $scope.names.length; ++i) {
            var name = $scope.names[i].name.trim();
            if (name !== '') {
                $scope.data.names.push(name);
            }
        }
        $rootScope.data = $scope.data;
        $scope.loading = true;
        if (getLatexCode) {
            url = url + "?tex";
        } else {
            newWindow = window.open()
        }
        $http.post(url, $scope.data).
            then(function (response) {
                $scope.loading = false;
                if (getLatexCode) {
                    console.log(response.data);
                    $rootScope.code = response.data.tex;
                    $location.path('/tex');
                } else {
                    $scope.url = newWindow.location = response.data.url;
                }
            }, function (response) {
                $scope.loading = false;
                console.log(response);
                // TODO: proper error message to the user
            });
    }
});

caratulaControllers.controller('CodeController', function ($scope, $rootScope) {
    $scope.code = $rootScope.code;
});
