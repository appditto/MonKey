<script>
  import axios from "axios";
  import { genAddress, validateAddress } from "../plugins/address.js";
  import {
    generatorIn,
    generatorOut,
    monkeyLoadingIn,
    monkeyContainerIn,
    monkeyContainerOut,
    formIn,
    formOut,
    curtainIn,
    againIn,
    againOut,
  } from "../plugins/transitions.js";
  export let showGenerator = false;
  let inputValue;
  let inputError = false;
  let inputFocused = false;
  let inputHovered = false;
  let monkeySvg;
  /* Variables for the generation animation/transitions */
  let monkeyLoading = false;
  let monkeyLoaded = false;
  ///////////////////////////////////////////
  async function getMonkey(address) {
    try {
      return axios.get(
        "https://testmonkey.appditto.com/api/v1/banano?address=" + address
      );
    } catch (e) {
      console.error(e);
    }
  }
  async function generateMonkey(address) {
    if (validateAddress(address)) {
      monkeyLoading = true;
      console.log("Monkey loading");
      let monkeyResult = await getMonkey(address);
      if (monkeyResult.data) {
        monkeyLoaded = true;
        setTimeout(() => {
          monkeySvg = monkeyResult.data;
          monkeyLoading = false;
        }, 150);
      }
    } else {
      inputError = true;
    }
  }
  function generateRandomMonkey() {
    let address = genAddress();
    generateMonkey(address);
    setTimeout(() => {
      inputValue = address;
    }, 200);
  }
  function resetGeneration() {
    monkeyLoading = false;
    monkeyLoaded = false;
    inputError = false;
    monkeySvg = null;
  }
</script>

<!-- HTML -->
{#if showGenerator}
  <div
    in:generatorIn
    out:generatorOut
    class="max-w-md max-h-md generator flex flex-col bg-white absolute top-0
    mt-8 overflow-hidden"
  >
    <!-- MonKey loading animation -->
    {#if monkeyLoading}
      <div
        in:monkeyLoadingIn={{ delay: 200 }}
        class="w-full h-full flex flex-row justify-center items-center absolute
        left-0 top-0"
      >
        <div class="w-24 h-24 relative">
          <div class="w-full h-full absolute cube cube-grayLight" />
          <div class="w-full h-full absolute cube cube-brown" />
          <div class="w-full h-full absolute cube cube-brownLight" />
          <div class="w-full h-full absolute cube cube-gray" />
        </div>
      </div>
    {/if}
    <!-- MonKey container -->
    {#if monkeyLoaded}
      <div
        in:monkeyContainerIn
        out:monkeyContainerOut
        class="w-full h-auto absolute left-0 top-0"
      >
        {#if monkeySvg}
          {@html monkeySvg}
        {/if}
      </div>
    {/if}
    {#if monkeyLoaded}
      <!-- Again Button -->
      <div
        in:againIn={{ delay: 400 }}
        out:againOut
        class="w-full flex flex-row justify-center absolute bottom-0"
      >
        <button
          on:click={resetGeneration}
          class="bg-primary btn-primary text-white text-lg font-bold rounded-lg
          border-2 border-black px-6 md:px-8 py-1 mx-4 md:mx-8 my-4 md:my-5"
        >
          Again!
        </button>
      </div>
    {/if}
    <!-- Input, Show Me & Randomize -->
    {#if !monkeyLoading && !monkeyLoaded}
      <div class="w-full h-full flex flex-col relative">
        {#if !monkeyLoading && !monkeyLoaded}
          <form
            out:formOut
            in:formIn={{ delay: 100 }}
            on:submit|preventDefault={generateMonkey(inputValue)}
            class="flex flex-col items-center my-auto relative mx-4 md:mx-6"
          >
            <div class="w-full">
              <label
                class="{inputError ? 'text-danger' : inputFocused || inputHovered ? 'text-brownLight' : 'text-gray'}
                absolute bg-white rounded-lg top-0 left-0 ml-4 -mt-4 px-2
                text-xl font-bold transition-all duration-200 ease-out"
                for="bananoAddress"
              >
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
                w-full text-xl font-bold px-4 py-3 border-3 rounded-xl
                transition-all duration-200 ease-out"
                type="text"
                autocomplete="off"
                placeholder="Enter your address"
              />
            </div>
            <button
              on:click={generateMonkey(inputValue)}
              class="w-full bg-primary btn-primary text-white text-xl font-bold
              rounded-xl border-black border-2 px-6 py-2 mx-auto mt-3"
            >
              Show Me
            </button>
          </form>
        {/if}
        {#if !monkeyLoading && !monkeyLoaded}
          <div
            out:formOut
            in:formIn={{ delay: 100 }}
            class="w-full flex flex-row justify-center absolute bottom-0"
          >
            <button
              on:click={generateRandomMonkey}
              class="bg-primary btn-primary text-white text-lg font-bold
              rounded-lg border-black border-2 px-6 md:px-8 py-1 my-4 md:my-5"
            >
              Randomize
            </button>
          </div>
        {/if}
      </div>
    {/if}
    <!-- Curtain -->
    {#if monkeyLoaded}
      <div
        in:curtainIn
        class="w-full h-full absolute transform -translate-y-full
        overflow-hidden"
      >
        <div class="w-full h-full bg-grayLight absolute" />
        <div class="w-full h-full bg-brownLight absolute mt-1/8" />
        <div class="w-full h-full bg-brown absolute mt-1/20" />
        <div class="w-full h-full bg-gray absolute mt-1/35" />
      </div>
    {/if}
  </div>
{/if}

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
    transform-origin: top center;
    overflow: hidden;
    animation: generatorAnimation;
    animation-duration: 2s;
    animation-iteration-count: infinite;
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
