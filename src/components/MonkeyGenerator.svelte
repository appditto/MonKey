<script>
  import axios from "axios";
  import { genAddress, validateAddress } from "../plugins/address.js";
  export let generatorVisibility = false;
  let inputValue;
  let inputError = false;
  let inputFocused = false;
  let inputHovered = false;
  let receivedMonkey = false;
  let monkeyLoading = false;
  let generationStarted = false;
  let showAgainButton = false;
  let monkeyContainer;
  let getMonkey = async (address) => {
    try {
      return axios.get(
        "https://natricon.com/api/v1/nano?svc=natricon.com&address=" + address
      );
    } catch (e) {
      console.error(e);
    }
  };
  let generateMonkey = async (address) => {
    if (validateAddress(address)) {
      generationStarted = true;
      setTimeout(() => {
        monkeyLoading = true;
      }, 125);
      let monkeyResult = await getMonkey(address);
      if (monkeyResult.data) {
        receivedMonkey = true;
        setTimeout(() => {
          monkeyLoading = false;
        }, 150);
        setTimeout(() => {
          monkeyContainer.innerHTML = monkeyResult.data;
        }, 250);
        setTimeout(() => {
          showAgainButton = true;
        }, 200);
      }
    } else {
      inputError = true;
    }
  };
  let resetGeneration = () => {
    monkeyContainer.innerHTML = "";
    receivedMonkey = false;
    monkeyLoading = false;
    generationStarted = false;
    showAgainButton = false;
    inputError = false;
  };
</script>

<style>
  .generator {
    width: calc(100vw - 2.5rem);
    height: calc(100vw - 2.5rem);
    border-radius: 1rem;
    border-width: 0rem;
    border-color: #404040;
    box-shadow: -0.5rem -0.5rem 0rem 0rem #404040,
      0.5rem -0.5rem 0rem 0rem #7f6145, 0.5rem 0.5rem 0rem 0rem #ffcd98,
      -0.5rem 0.5rem 0rem 0rem #9b9ba1;
    transition: opacity 0.2s cubic-bezier(0.215, 0.51, 0.355, 1),
      transform 0.5s cubic-bezier(0.215, 0.51, 0.355, 1);
    transform-origin: top center;
    overflow: hidden;
    transform: scale(1);
    opacity: 1;
    animation: generatorAnimation;
    animation-duration: 2s;
    animation-iteration-count: infinite;
  }
  .closed {
    transform: scale(0.5) !important;
    opacity: 0 !important;
  }
  @media screen and (min-width: 768px) {
    .generator {
      width: 50vw;
      height: 50vw;
      max-width: 50vw;
      max-height: 50vw;
    }
  }
  @media screen and (min-width: 768px) {
    .generator {
      min-width: 24rem;
      min-height: 24rem;
      width: 20vw;
      height: 20vw;
      max-width: 20vw;
      max-height: 20vw;
    }
  }
  @keyframes generatorAnimation {
    0% {
      box-shadow: -0.5rem -0.5rem 0rem 0rem #404040,
        0.5rem -0.5rem 0rem 0rem #7f6145, 0.5rem 0.5rem 0rem 0rem #ffcd98,
        -0.5rem 0.5rem 0rem 0rem #9b9ba1;
    }
    25% {
      box-shadow: 0.5rem -0.5rem 0rem 0rem #404040,
        0.5rem 0.5rem 0rem 0rem #7f6145, -0.5rem 0.5rem 0rem 0rem #ffcd98,
        -0.5rem -0.5rem 0rem 0rem #9b9ba1;
    }
    50% {
      box-shadow: 0.5rem 0.5rem 0rem 0rem #404040,
        -0.5rem 0.5rem 0rem 0rem #7f6145, -0.5rem -0.5rem 0rem 0rem #ffcd98,
        0.5rem -0.5rem 0rem 0rem #9b9ba1;
    }
    75% {
      box-shadow: -0.5rem 0.5rem 0rem 0rem #404040,
        -0.5rem -0.5rem 0rem 0rem #7f6145, 0.5rem -0.5rem 0rem 0rem #ffcd98,
        0.5rem 0.5rem 0rem 0rem #9b9ba1;
    }
    100% {
      box-shadow: -0.5rem -0.5rem 0rem 0rem #404040,
        0.5rem -0.5rem 0rem 0rem #7f6145, 0.5rem 0.5rem 0rem 0rem #ffcd98,
        -0.5rem 0.5rem 0rem 0rem #9b9ba1;
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
    transition: all 0.58s;
  }
  .curtain-2 {
    transition: all 0.59s;
  }
  .curtain-3 {
    transition: all 0.68s;
  }
  .curtain-4 {
    transition: all 0.75s;
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
  class="max-w-md max-h-md generator {!generatorVisibility ? 'closed' : ''} flex
  flex-col bg-white absolute top-0 mt-8 overflow-hidden">
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
  {#if receivedMonkey}
    <!-- Again Button -->
    <div class="w-full flex flex-row justify-center absolute bottom-0">
      <button
        disabled={!showAgainButton}
        on:click={() => {
          resetGeneration();
        }}
        class="{showAgainButton ? 'scale-100 opacity-100' : 'scale-0 opacity-50'}
        transform duration-200 ease-out bg-primary btn-primary text-white
        text-lg font-bold rounded-lg border-2 border-black px-6 md:px-8 py-1
        mx-4 md:mx-8 my-4 md:my-5">
        Again!
      </button>
    </div>
  {/if}
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
    transform duration-200 ease-out w-full h-full flex flex-col relative">
    <form
      on:submit|preventDefault={() => {
        generateMonkey(inputValue);
      }}
      class="flex flex-col items-center my-auto relative mx-4 md:mx-6">
      <div class="w-full">
        <label
          class="{inputError ? 'text-danger' : inputFocused || inputHovered ? 'text-brownLight' : 'text-gray'}
          absolute bg-white rounded-lg top-0 left-0 ml-4 -mt-4 px-2 text-xl
          font-bold transition-all duration-200 ease-out"
          for="bananoAddress">
          Address
        </label>
        <input
          name="bananoAddress"
          id="bananoAddress"
          on:blur={() => {
            inputFocused = false;
          }}
          on:focus={() => {
            inputFocused = true;
          }}
          on:mouseenter={() => {
            inputHovered = true;
          }}
          on:mouseleave={() => {
            inputHovered = false;
          }}
          bind:value={inputValue}
          on:input={() => {
            if (inputError) {
              inputError = false;
            }
          }}
          class="{inputError ? 'border-danger text-danger' : 'text-gray border-primary focus:border-brownLight hover:border-brownLight'}
          w-full text-xl font-bold px-4 py-3 border-3 rounded-xl transition-all
          duration-200 ease-out"
          type="text"
          autocomplete="off"
          placeholder="Enter your address" />
      </div>
      <button
        disabled={generationStarted}
        on:click={() => {
          generateMonkey(inputValue);
        }}
        class="w-full bg-primary btn-primary text-white text-xl font-bold
        rounded-xl border-2 border-black px-6 py-2 mx-auto mt-3">
        Show Me
      </button>
    </form>
    <div class="w-full flex flex-row justify-center absolute bottom-0">
      <button
        disabled={generationStarted}
        on:click={() => {
          let address = genAddress();
          generateMonkey(address);
          setTimeout(() => {
            inputValue = address;
          }, 200);
        }}
        class="bg-primary btn-primary text-white text-lg font-bold rounded-lg
        border-2 border-black px-6 md:px-8 py-1 my-4 md:my-5">
        Randomize
      </button>
    </div>
  </div>
</div>
