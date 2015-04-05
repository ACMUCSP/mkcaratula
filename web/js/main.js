angular.module('CaratulaApp', [])
    .controller('CaratulaController', function ($scope, $http) {
        $scope.data = {
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
            if (item.id === $scope.names.length && $scope.names.length < 6) {
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
            console.log($scope.data);
        }
    });