<script>
  import axios from "axios";
  export let generatorVisibility = false;
  let inputValue;
  let receivedMonkey = false;
  let monkeyLoading = false;
  let showAgainButton = false;
  let monkeyContainer;
  let getMonkey = async () => {
    try {
      return axios.get(
        "https://natricon.com/api/v1/nano?svc=natricon.com&address=" +
          inputValue
      );
    } catch (e) {
      console.error(e);
    }
  };
  let generateMonkey = async () => {
    monkeyLoading = true;
    let monkeyResult = await getMonkey();
    if (monkeyResult.data) {
      receivedMonkey = true;
      console.log(monkeyResult);
      monkeyContainer.innerHTML = monkeyResult.data;
      monkeyLoading = false;
      showAgainButton = true;
    }
  };
</script>

<style>
  .generator {
    width: calc(100vw - 1rem);
    height: calc(100vw - 1rem);
    border-radius: 1rem;
    border-width: 0.25rem;
    box-shadow: 0rem 0.6rem 0rem 0rem #404040;
    transition: opacity 0.2s cubic-bezier(0.215, 0.61, 0.355, 1),
      transform 0.4s cubic-bezier(0.215, 0.61, 0.355, 1);
    transform-origin: top center;
    overflow: hidden;
    transform: scale(1);
    opacity: 1;
  }
  .closed {
    transform: scale(0.5) !important;
    opacity: 0 !important;
  }
  @media screen and (min-width: 768px) {
    .generator {
      width: 50vw;
      height: 50vw;
    }
  }
  @media screen and (min-width: 768px) {
    .generator {
      min-width: 24rem;
      min-height: 24rem;
      width: 20vw;
      height: 20vw;
    }
  }
</style>

<!-- HTML -->
<div
  class="generator {!generatorVisibility ? 'closed' : ''} border-primary flex
  flex-col bg-white absolute top-0 mt-5">
  <div bind:this={monkeyContainer} />
  <!-- Input, Show Me & Randomize -->
  <div
    class="{monkeyLoading || receivedMonkey ? 'scale-0' : 'scale-100'} transform
    duration-200 ease-out w-full h-full flex flex-col px-4 md:px-8 py-4 md:py-5">
    <div class="flex flex-col items-center my-auto">
      <input
        bind:value={inputValue}
        class="w-full text-xl font-bold px-4 py-2 border-2 border-black
        rounded-xl"
        type="text"
        placeholder="enter your address" />
      <button
        disabled={monkeyLoading || receivedMonkey}
        on:click={() => {
          generateMonkey();
        }}
        class="w-full bg-primary btn-primary text-white text-xl font-bold
        rounded-xl border-2 border-black px-6 py-2 mx-auto mt-3">
        Show Me
      </button>
    </div>
    <button
      disabled={monkeyLoading || receivedMonkey}
      class="bg-primary btn-primary text-white text-lg font-bold rounded-lg
      border-2 border-black px-6 md:px-8 py-1 mx-auto">
      Randomize
    </button>
  </div>
</div>
