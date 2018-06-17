/**
 * Created by jpierre on 2/23/16.
 */
$(function () {

});

var buildURL = "/api/builds";
var commitUrlBase = "https://github.com/nucklehead/patecho-seve/commit/";
var buildUrlBase = "https://travis-ci.com/nucklehead/patecho-seve/builds/";

var phonecatApp = angular.module('builds-app', []);


phonecatApp.controller('buildsCtrl', function ($scope, $sce, $http) {

    $http.get(buildURL)
    .then(function (reponse) {
        $scope.builds = reponse.data;
        $scope.buildUrl = $sce.trustAsResourceUrl("builds/" + $scope.builds[0].name + "/index.html");
        $scope.selectedIndex = 0;
        console.log($scope.builds);
    },
  function(error){
    console.log(error);
  });


    $scope.showResults = function (build, index) {
        $scope.selectedIndex = index;
        if (build.id === "denye") {
            $scope.buildUrl = $sce.trustAsResourceUrl("builds/denye/index.html");
        }
        else {
            $scope.buildUrl = $sce.trustAsResourceUrl("builds/" + build.name + "/index.html");
        }

    };

    $scope.getBuildURL = function (build, index) {
        $scope.selectedIndex = index;
        if (build.id === "denye") {
            return $sce.trustAsResourceUrl("#");
        }
        else {
            return $sce.trustAsResourceUrl(buildUrlBase + build.id );
        }

    };

    $scope.getCommitURL = function (build, index) {
        $scope.selectedIndex = index;
        return $sce.trustAsResourceUrl(commitUrlBase + build.commit_id );
    };
});
