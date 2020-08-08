<script>
  import { fade } from "svelte/transition";
  export let height = 0;
  export let offset = -750;
  export let resetHeightDelay = 0;
  export let onload = null;
  let className = "";
  export { className as class };
  let rootClass = "svelte-lazy" + (className ? " " + className : "");
  const transitionClass = "svelte-lazy-transition";
  const placeholderClass = "placeholder";
  let loaded = false;
  function load(node) {
    setHeight(node);
    const loadHandler = throttle((e) => {
      const top = node.getBoundingClientRect().top;
      const expectedTop = getExpectedTop(e, offset);
      if (top <= expectedTop) {
        loaded = true;
        resetHeight(node);
        onload && onload(node);
        removeListeners();
      }
    }, 200);
    loadHandler();
    addListeners();
    function addListeners() {
      document.addEventListener("scroll", loadHandler, true);
      window.addEventListener("resize", loadHandler);
    }
    function removeListeners() {
      document.removeEventListener("scroll", loadHandler, true);
      window.removeEventListener("resize", loadHandler);
    }
    return {
      destroy: () => {
        removeListeners();
      },
    };
  }
  function setHeight(node) {
    if (height) {
      node.style.height = typeof height === "number" ? height + "px" : height;
    }
  }
  function resetHeight(node) {
    // Add delay for remote resources like images to load
    setTimeout(() => {
      const img = node.querySelector("img");
      if (img && !img.complete) {
        node.addEventListener(
          "load",
          () => {
            node.style.height = "auto";
          },
          { capture: true, once: true }
        );
      } else {
        node.style.height = "auto";
      }
    }, resetHeightDelay);
  }
  function getExpectedTop(e, offset) {
    const height = getContainerHeight(e);
    return height + offset;
  }
  function getContainerHeight(e) {
    if (e && e.target && e.target.getBoundingClientRect) {
      return e.target.getBoundingClientRect().bottom;
    } else {
      return window.innerHeight;
    }
  }
  // From underscore souce code
  function throttle(func, wait, options) {
    let context, args, result;
    let timeout = null;
    let previous = 0;
    if (!options) options = {};
    const later = function () {
      previous = options.leading === false ? 0 : new Date();
      timeout = null;
      result = func.apply(context, args);
      if (!timeout) context = args = null;
    };
    return function (event) {
      const now = new Date();
      if (!previous && options.leading === false) previous = now;
      const remaining = wait - (now - previous);
      context = this;
      args = arguments;
      if (remaining <= 0 || remaining > wait) {
        if (timeout) {
          clearTimeout(timeout);
          timeout = null;
        }
        previous = now;
        result = func.apply(context, args);
        if (!timeout) context = args = null;
      } else if (!timeout && options.trailing !== false) {
        timeout = setTimeout(later, remaining);
      }
      return result;
    };
  }
</script>

<div use:load class="w-full relative {rootClass}">
  <div
    class="{transitionClass}
    {loaded ? 'opacity-0' : 'opacity-100'} w-full transition-opacity
    duration-500">
    <slot name="placeholder" />
  </div>
  <div
    class="{transitionClass}
    {loaded ? 'opacity-100' : 'opacity-0'} w-full transition-opacity
    duration-300 ease-out absolute top-0 left-0">
    <slot name="content" />
  </div>
</div>
