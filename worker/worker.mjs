import "./polyfill_performance.js";
import "./wasm_exec.js";
import mod from "../dist/app.wasm";

const go = new Go();

export default {
  async fetch(req, env, ctx) {
    const readyPromise = new Promise((resolve) => {
      globalThis.ready = resolve;
    });
    const instance = await WebAssembly.instantiate(mod, go.importObject);
    go.run(instance);
    await readyPromise;
    return handleRequest(req, { env, ctx });
  },
};
