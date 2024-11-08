import Prism from "prismjs";

(function () {
  if (typeof Prism === "undefined" || typeof document === "undefined") {
    return;
  }

  Prism.plugins.lineInsert = {
    wrapLines: function (env) {
      if (!env || !env.element) return;

      console.log(env.element.innerHTML);

      const lines = env.element.innerHTML.split("\n");
      console.log(lines);

      env.element.innerHTML = lines
        .map((line) => `<div class="line-wrapper">${line}</div>`)
        .join("");
    },
  };

  Prism.hooks.add("line-numbers", function (env) {
    Prism.plugins.lineInsert.wrapLines(env);
  });
})();
