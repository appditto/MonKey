<script>
  import axios from "axios";
  export let generatorVisibility = false;
  let inputValue;
  let receivedMonkey = false;
  let monkeyLoading = false;
  let generationStarted = false;
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
    generationStarted = true;
    setTimeout(() => {
      monkeyLoading = true;
    }, 100);
    let monkeyResult = await getMonkey();
    if (monkeyResult.data) {
      receivedMonkey = true;
      setTimeout(() => {
        monkeyLoading = false;
      }, 150);
      setTimeout(() => {
        monkeyContainer.innerHTML = monkeyResult.data;
      }, 250);
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
  .cube {
    border-radius: 15%;
    border-width: 3px;
    transform: translate(0rem, 0rem);
  }
  .cube-brownLight {
    background-color: #ffcd98;
    border-color: #cd9e6c;
    box-shadow: 0rem 0.3rem 0rem 0rem #cd9e6c;
    animation-name: animation-1;
    animation-duration: 1.3s;
    animation-iteration-count: infinite;
    animation-delay: -0.5s;
  }
  .cube-brown {
    background-color: #7f6145;
    border-color: #6c4725;
    box-shadow: 0rem 0.3rem 0rem 0rem #6c4725;
    animation-name: animation-2;
    animation-duration: 1.1s;
    animation-iteration-count: infinite;
    animation-delay: -0.25s;
  }
  .cube-gray {
    background-color: #404040;
    border-color: #000000;
    box-shadow: 0rem 0.3rem 0rem 0rem #000000;
    animation-name: gray;
    animation-name: animation-2;
    animation-duration: 1s;
    animation-iteration-count: infinite;
    animation-delay: -0.75s;
  }
  .cube-grayLight {
    background-color: #9b9ba1;
    border-color: #72727a;
    box-shadow: 0rem 0.3rem 0rem 0rem #72727a;
    animation-name: animation-1;
    animation-duration: 1.2s;
    animation-iteration-count: infinite;
  }
  .hide-curtain {
    transform: translateY(-100%);
  }
  .show-curtain {
    transform: translateY(100%);
  }
  .curtain-1 {
    transition: all 0.5s;
  }
  .curtain-2 {
    transition: all 0.6s;
  }
  .curtain-3 {
    transition: all 0.7s;
  }
  .curtain-4 {
    transition: all 0.8s;
  }
  @keyframes animation-1 {
    0% {
      transform: translate(-2rem, 2rem);
    }
    50% {
      transform: translate(2rem, -2rem);
    }
    100% {
      transform: translate(-2rem, 2rem);
    }
  }
  @keyframes animation-2 {
    0% {
      transform: translate(2.5rem, 2.5rem);
    }
    50% {
      transform: translate(-2.5rem, -2.5rem);
    }
    100% {
      transform: translate(2.5rem, 2.5rem);
    }
  }
</style>

<!-- HTML -->
<div
  class="generator {!generatorVisibility ? 'closed' : ''} border-primary flex
  flex-col bg-white absolute top-0 mt-5 overflow-hidden">
  <!-- MonKey loading animation -->
  {#if generationStarted}
    <div
      class="{monkeyLoading ? 'scale-100 opacity-100' : 'scale-0 opacity-50'}
      transform duration-200 ease-out w-full h-full flex flex-row justify-center
      items-center absolute left-0 top-0">
      <div class="w-24 h-24 relative">
        <div class="w-full h-full absolute cube cube-grayLight" />
        <div class="w-full h-full absolute cube cube-brown" />
        <div class="w-full h-full absolute cube cube-brownLight" />
        <div class="w-full h-full absolute cube cube-gray" />
      </div>
    </div>
  {/if}
  <!-- MonKey container -->
  <div
    bind:this={monkeyContainer}
    class="{receivedMonkey ? 'translate-y-0' : '-translate-y-20'} transform
    duration-700 ease-out" />
  <!-- Curtain -->
  {#if generationStarted}
    <div
      class="{receivedMonkey ? 'show-curtain' : 'hide-curtain'} curtain-4 w-full
      h-full bg-grayLight absolute" />
    <div
      class="{receivedMonkey ? 'show-curtain' : 'hide-curtain'} curtain-3 w-full
      h-full bg-brownLight absolute" />
    <div
      class="{receivedMonkey ? 'show-curtain' : 'hide-curtain'} curtain-2 w-full
      h-full bg-brown absolute" />
    <div
      class="{receivedMonkey ? 'show-curtain' : 'hide-curtain'} curtain-1 w-full
      h-full bg-gray absolute" />
  {/if}
  <!-- Input, Show Me & Randomize -->
  <div
    class="{generationStarted ? 'scale-0 opacity-50' : 'scale-100 opacity-100'}
    transform duration-200 ease-out w-full h-full flex flex-col px-4 md:px-8
    py-4 md:py-5">
    <div class="flex flex-col items-center my-auto">
      <input
        bind:value={inputValue}
        class="w-full text-xl font-bold px-4 py-2 border-2 border-black
        rounded-xl"
        type="text"
        placeholder="enter your address" />
      <button
        disabled={generationStarted}
        on:click={() => {
          generateMonkey();
        }}
        class="w-full bg-primary btn-primary text-white text-xl font-bold
        rounded-xl border-2 border-black px-6 py-2 mx-auto mt-3">
        Show Me
      </button>
    </div>
    <button
      disabled={generationStarted}
      class="bg-primary btn-primary text-white text-lg font-bold rounded-lg
      border-2 border-black px-6 md:px-8 py-1 mx-auto">
      Randomize
    </button>
  </div>
</div>
