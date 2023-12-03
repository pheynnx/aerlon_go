import { build } from "esbuild";
// import glob from "tiny-glob";
import { getGlobals } from "common-es";

(async () => {
  // let entryPoints = await glob("./web/source/js/scripts/*.{js,ts}");

  const { __dirname, __filename } = getGlobals(import.meta.url);

  const entries = [];

  let entryPoints = entries.map((e) => `${__dirname}\\scripts\\${e}`);
  await build({
    entryPoints,
    entryNames: "[name]-compiled",
    bundle: true,
    minify: true,
    sourcemap: false,
    outdir: "web/static/js",
  });
})();
