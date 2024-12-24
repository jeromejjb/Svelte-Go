import svelte from 'rollup-plugin-svelte';  // Import the svelte plugin
import resolve from '@rollup/plugin-node-resolve';  // For resolving node modules
import commonjs from '@rollup/plugin-commonjs';  // To convert CommonJS to ES6
import { terser } from '@rollup/plugin-terser';  // Correct terser plugin for Rollup 3.x
import css from 'rollup-plugin-css-only';  // For handling CSS output

const production = process.env.NODE_ENV === 'production';

export default {
  input: 'src/main.js',  // Your entry file
  output: {
    file: 'public/build/bundle.js',  // Output file
    format: 'iife',  // Immediately Invoked Function Expression
    name: 'app',  // The global variable name for the bundle
    sourcemap: true,  // Generate sourcemaps for debugging
  },
  plugins: [
    svelte({
      dev: !production,  // Enable development mode for easier debugging
      css: css => {
        css.write('public/build/bundle.css');  // Write the CSS to a separate file
      },
    }),
    resolve(),
    commonjs(),
    css({ output: 'public/build/styles.css' }),  // Output CSS file

    // Minify the output in production mode
    production && terser(),
  ],
};
