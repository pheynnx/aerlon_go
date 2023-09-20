import { build } from "esbuild";

await build({
  entryPoints: [],
  entryNames: "[name]-compiled",
  bundle: true,
  minify: true,
  sourcemap: false,
  outdir: "web/static/js",
}).catch(() => process.exit(1));
