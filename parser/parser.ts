import { Marked } from "marked";
import { markedHighlight } from "marked-highlight";
import { getHighlighter, bundledLanguages } from "shikiji";

const shiki = await getHighlighter({
  themes: ["material-theme-palenight", "min-light"],
  langs: [...Object.keys(bundledLanguages)],
});

const marked = new Marked(
  markedHighlight({
    highlight(code, lang) {
      try {
        const cody = shiki.codeToHtml(code, {
          lang,
          themes: {
            light: "min-light",
            dark: "material-theme-palenight",
          },
          cssVariablePrefix: "--theme-",
          defaultColor: "",
          transformers: [
            {
              pre(hast) {
                return {
                  type: "element",
                  tagName: "pre",
                  properties: {
                    class: "shiki",
                  },
                  children: hast.children,
                };
              },
            },
          ],
        });

        return cody;
      } catch (error) {
        return;
      }
    },
  }),
  {
    renderer: {
      code(code, lang, escaped) {
        return `<div class="code-block"><span class="language-name">${lang}</span>${escaped ? code : code
          }</div>`;
      },
    },
  }
);

const message = process.argv[2];
try {
  let result = await marked.parse(message)

  process.stdout.write(result);
} catch (error) {
  process.stderr.write(`${error}`)
}