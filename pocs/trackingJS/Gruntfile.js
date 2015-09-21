/*  */
'use strict';

module.exports = function(grunt) {
var SERVER_PORT = 9000;


    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        connect: {
            example: {
                port: SERVER_PORT,
                base: 'public'
            }
        }
    });

    grunt.loadNpmTasks('grunt-connect');
    grunt.registerTask('default', 'connect:example');

};
