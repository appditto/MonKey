<script>
  import axios from "axios";
  import { genAddress, validateAddress } from "../plugins/address.js";
  import { fadeAndScaleIn, fadeAndScaleOut } from "../plugins/transitions.js";
  export let showGenerator = false;
  let inputValue;
  let inputError = false;
  let inputFocused = false;
  let inputHovered = false;
  let monkeyContainer;
  /* Variables for the generation animation */
  let hideForm = false;
  let hideFormAnimation = false;
  let showLoading = false;
  let showLoadingAnimation = false;
  let showCurtain = false;
  let showCurtainAnimation = false;
  let showMonkeyContainer = false;
  let showMonkeyContainerAnimation = false;
  let showAgainButton = false;
  let showAgainButtonAnimation = false;
  let hideMonkeyContainer = false;
  let toHideMonkeyContainer = false;
  /* ////////////////////////////////////// */
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
      hideFormAnimation = true;
      setTimeout(() => {
        hideForm = true;
      }, 175);
      showLoading = true;
      setTimeout(() => {
        showLoadingAnimation = true;
      }, 175);
      let monkeyResult = await getMonkey(address);
      if (monkeyResult.data) {
        showMonkeyContainer = true;
        showCurtain = true;
        setTimeout(() => {
          showCurtainAnimation = true;
          showMonkeyContainerAnimation = true;
        }, 25);
        setTimeout(() => {
          monkeyContainer.innerHTML = monkeyResult.data;
          showLoading = false;
        }, 200);
        showAgainButton = true;
        setTimeout(() => {
          showAgainButtonAnimation = true;
        }, 400);
        setTimeout(() => {
          showCurtain = false;
        }, 725);
      }
    } else {
      inputError = true;
    }
  };
  let resetGeneration = () => {
    showLoading = false;
    showLoadingAnimation = false;
    showCurtain = false;
    showCurtainAnimation = false;
    toHideMonkeyContainer = true;
    showAgainButtonAnimation = false;
    setTimeout(() => {
      showAgainButton = false;
    }, 300);
    hideForm = false;
    inputError = false;
    setTimeout(() => {
      hideMonkeyContainer = true;
    }, 25);
    setTimeout(() => {
      showMonkeyContainer = false;
      showMonkeyContainerAnimation = false;
      hideMonkeyContainer = false;
      toHideMonkeyContainer = false;
    }, 450);
    setTimeout(() => {
      hideFormAnimation = false;
    }, 125);
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
  .hide-curtain {
    transform: translateY(-100%);
  }
  .show-curtain {
    transform: translateY(100%);
  }
  .curtain-1 {
    transition: all 0.49s;
  }
  .curtain-2 {
    transition: all 0.55s;
  }
  .curtain-3 {
    transition: all 0.62s;
  }
  .curtain-4 {
    transition: all 0.7s;
  }
  .monkey-container {
    transition: all 0.55s ease-out;
  }
  .hidden-monkey-container {
    transform: translateY(-20%);
  }
  .show-monkey-container {
    transform: translateY(0%);
  }
  .to-hide-monkey-container {
    transition: all 0.4s ease-out;
    opacity: 1;
  }
  .hide-monkey-container {
    transform: translateY(-100%);
    opacity: 0;
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
{#if showGenerator}
  <div
    in:fadeAndScaleIn
    out:fadeAndScaleOut
    class="max-w-md max-h-md generator flex flex-col bg-white absolute top-0
    mt-8 overflow-hidden">
    <!-- MonKey loading animation -->
    {#if showLoading}
      <div
        class="{showLoadingAnimation ? 'scale-100 opacity-100' : 'scale-0 opacity-50'}
        transform duration-200 ease-out w-full h-full flex flex-row
        justify-center items-center absolute left-0 top-0">
        <div class="w-24 h-24 relative">
          <div class="w-full h-full absolute cube cube-grayLight" />
          <div class="w-full h-full absolute cube cube-brown" />
          <div class="w-full h-full absolute cube cube-brownLight" />
          <div class="w-full h-full absolute cube cube-gray" />
        </div>
      </div>
    {/if}
    <!-- MonKey container -->
    {#if showMonkeyContainer}
      <div
        bind:this={monkeyContainer}
        class="{showMonkeyContainerAnimation ? (hideMonkeyContainer ? 'hide-monkey-container' : 'show-monkey-container') : 'hidden-monkey-container'}
        {toHideMonkeyContainer ? 'to-hide-monkey-container' : 'monkey-container'}
        w-full h-auto absolute left-0 top-0" />
    {/if}
    {#if showAgainButton}
      <!-- Again Button -->
      <div class="w-full flex flex-row justify-center absolute bottom-0">
        <button
          disabled={!showAgainButtonAnimation}
          on:click={() => {
            resetGeneration();
          }}
          class="{showAgainButtonAnimation ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-20'}
          transform duration-350 ease-out bg-primary btn-primary text-white
          text-lg font-bold rounded-lg border-2 border-black px-6 md:px-8 py-1
          mx-4 md:mx-8 my-4 md:my-5">
          Again!
        </button>
      </div>
    {/if}
    <!-- Curtain -->
    {#if showCurtain}
      <div
        class="{showCurtainAnimation ? 'show-curtain' : 'hide-curtain'}
        curtain-4 w-full h-full bg-grayLight absolute" />
      <div
        class="{showCurtainAnimation ? 'show-curtain' : 'hide-curtain'}
        curtain-3 w-full h-full bg-brownLight absolute" />
      <div
        class="{showCurtainAnimation ? 'show-curtain' : 'hide-curtain'}
        curtain-2 w-full h-full bg-brown absolute" />
      <div
        class="{showCurtainAnimation ? 'show-curtain' : 'hide-curtain'}
        curtain-1 w-full h-full bg-gray absolute" />
    {/if}
    <!-- Input, Show Me & Randomize -->
    {#if !hideForm}
      <div class="w-full h-full flex flex-col relative">
        <form
          on:submit|preventDefault={() => {
            generateMonkey(inputValue);
          }}
          class="{hideFormAnimation ? 'scale-0 opacity-25' : 'scale-100 opacity-100'}
          transform duration-200 ease-out flex flex-col items-center my-auto
          relative mx-4 md:mx-6">
          <div class="w-full">
            <label
              class="{inputError ? 'text-danger' : inputFocused || inputHovered ? 'text-brownLight' : 'text-gray'}
              absolute bg-white rounded-lg top-0 left-0 ml-4 -mt-4 px-2 text-xl
              font-bold transition-all duration-200 ease-out"
              for="bananoAddress">
              Address
            </label>
            <input
              disabled={hideFormAnimation}
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
              placeholder="Enter your address" />
          </div>
          <button
            disabled={hideFormAnimation}
            on:click={() => {
              generateMonkey(inputValue);
            }}
            class="w-full bg-primary btn-primary text-white text-xl font-bold
            rounded-xl border-black border-2 px-6 py-2 mx-auto mt-3">
            Show Me
          </button>
        </form>
        <div
          class="{hideFormAnimation ? 'scale-0 opacity-25' : 'scale-100 opacity-100'}
          transform duration-200 ease-out w-full flex flex-row justify-center
          absolute bottom-0">
          <button
            disabled={hideFormAnimation}
            on:click={() => {
              let address = genAddress();
              generateMonkey(address);
              setTimeout(() => {
                inputValue = address;
              }, 200);
            }}
            class="bg-primary btn-primary text-white text-lg font-bold
            rounded-lg border-black border-2 px-6 md:px-8 py-1 my-4 md:my-5">
            Randomize
          </button>
        </div>
      </div>
    {/if}
  </div>
{/if}
