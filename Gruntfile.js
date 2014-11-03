module.exports = function(grunt) {
  grunt.initConfig({
    copy: {
      main: {
        files: [
          {
            expand: true,
            flatten: true,
            src: [
              './bower_components/angular/**.min.js',
              './bower_components/angular/**.min.js.map',
              './bower_components/angular-resource/**.min.js',
              './bower_components/angular-resource/**.min.js.map',
              './bower_components/angular-route/**.min.js',
              './bower_components/angular-route/**.min.js.map'
            ],
            dest: './static/js/',
            filter: 'isFile',
          }
        ]
      }
    },
    less: {
      development: {
        options: {
          paths: ["./less"],
          compress: true,
          reload: true
        },
        files: {
          "./static/css/main.css": "./templates/main.less"
        }
      }
    },
    uglify: {
      development: {
        options: {
          paths: ["./javascript"],
          compress: true,
          reload: true
        },
        files: {}
      },
    },
    watch: {
      less: {
        files: ["./templates/**/*.less"],
        tasks: ["less"]
      },
      uglify: {
        files: ["./src/js/**/*.js"],
        tasks: ["uglify"]                
      },
    },
  });
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-less');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.registerTask('build', ['copy', 'less', 'uglify']);
  grunt.registerTask('heroku', 'build');
  grunt.registerTask('heroku:', 'build');
  grunt.registerTask('heroku:development', 'build');
  grunt.registerTask('heroku:production', 'build');
};
