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
        cat: 'inmasc',
        name: '',
        sem: '1er',
        tex: false
    };
    $scope.names = [{id: 1, name: ''}, {id: 2, name: ''}];
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
        if ($scope.names.length <= 2) {
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
        var previousName = $scope.data.name;
        if ($scope.form.$invalid) {
            return;
        }
        $scope.data.tex = getLatexCode;
        if ($scope.data.cat === 'grupal') {
            var names = $scope.names[0].name;
            for (var i = 1; i < $scope.names.length; ++i) {
                if ($scope.names[i].name !== '') {
                    names += '/' + $scope.names[i].name;
                }
            }
            $scope.data.name = names;
        }
        $rootScope.data = $scope.data;
        $scope.loading = true;
        $http.post(genUrl, $scope.data).
            success(function (data, status, headers, config) {
                $scope.loading = false;
                if (status === 201) {
                    var newWindow = window.open(headers('Location'), '_blank');
                    $scope.url = data;
                } else {
                    $rootScope.code = data;
                    $location.path('/tex');
                }
                $scope.data.name = previousName;
            }).
            error(function (data, status, headers, config) {
                $scope.loading = false;
                if (status === 400) {
                    $rootScope.code = data;
                    $location.path('/error');
                } else {
                    // inform to dev
                }
                $scope.data.name = previousName;
            });
    }
});

caratulaControllers.controller('CodeController', function ($scope, $rootScope) {
    $scope.code = $rootScope.code;
});
