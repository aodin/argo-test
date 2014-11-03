module.exports = function(grunt) {
  grunt.initConfig({
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
      copy: {
        files: {}
      }
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
  grunt.loadNpmTasks('grunt-contrib-less');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.registerTask('build', ['less', 'uglify']);
  grunt.registerTask('heroku', 'build');
  grunt.registerTask('heroku:', 'build');
  grunt.registerTask('heroku:development', 'build');
  grunt.registerTask('heroku:production', 'build');
};
