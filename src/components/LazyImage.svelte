<script>
  export let offset = 200;
  export let onload = null;
  export let classes;
  let loaded = false;
  function load(node) {
    const loadHandler = throttle((e) => {
      const top = node.getBoundingClientRect().top;
      const expectedTop = getExpectedTop(e, offset);
      if (top <= expectedTop) {
        loaded = true;
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

<div class={classes}>
  <div use:load class="w-full relative">
    <div
      class="{loaded ? 'opacity-0' : 'opacity-100'} w-full transition-opacity
      duration-500">
      <slot name="placeholder" />
    </div>
    <div
      class="{loaded ? 'opacity-100' : 'opacity-0'} w-full transition-opacity
      duration-300 ease-out absolute top-0 left-0">
      <slot name="content" />
    </div>
  </div>
</div>
