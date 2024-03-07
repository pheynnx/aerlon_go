import { build } from "esbuild";

(async () => {
  await build({
    entryPoints: ["./parser/parser.ts"],
    entryNames: "[name]-compiled",
    bundle: true,
    minify: true,
    sourcemap: false,
    splitting: false,
    format: "esm",
    outdir: "./parser",
  });
})();
