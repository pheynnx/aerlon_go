import { build } from "esbuild";

await build({
  entryPoints: [
    "web/source/js/scripts/navigation.js",
    "web/source/js/scripts/index.js",
  ],
  entryNames: "[name]-compiled",
  bundle: true,
  minify: true,
  sourcemap: false,
  outdir: "web/static/js",
}).catch(() => process.exit(1));
