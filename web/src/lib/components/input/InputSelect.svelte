<script lang="ts">
  import InputOption from '$lib/components/input/InputOption.svelte';
  import Response from '$lib/components/input/Response.svelte';
  import { fade, slide } from 'svelte/transition';

  type Props = {
    label: string;
    options: string[];
    placeholder?: string;
    responseMessage?: string;
    submitOption: (val: string) => Promise<InputError>;
  };
  let { label, options, responseMessage = '', placeholder = '', submitOption }: Props = $props();

  let selectedOption: string = $state(placeholder);
  let showOptions: boolean = $state(false);

  let optionsIndex: number = $state(0);

  // response is a promise so #await can be used to handle pending state
  let response: Promise<InputError> = $state(
    Promise.resolve({
      error: false,
      message: responseMessage
    })
  );

  function clampOptionIndex(n: number) {
    optionsIndex = n < 0 ? 0 : n >= options.length ? options.length - 1 : n;
  }
</script>

<div class="flex items-center gap-2">
  <div class="relative flex h-12 w-80 items-center">
    <label
      for={label}
      class="relative mt-2 w-full border-2 border-jfgray-700 text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 hover:border-ctp-lavender-50/50 {selectedOption
        ? 'bg-jfgray-900'
        : 'bg-jfgray-800'}"
    >
      <span class="absolute -top-1 left-2 bg-jfgray-800 px-1 text-base leading-1">{label}</span>
      <div class="relative flex h-10">
        <input
          type="button"
          id={label}
          onmousedown={() => {
            showOptions = true;
          }}
          onfocusout={() => {
            showOptions = false;
          }}
          onkeydown={(e) => {
            const key = e.key;
            if (key === 'ArrowDown') {
              e.preventDefault();
              clampOptionIndex(optionsIndex + 1);
              return;
            }
            if (key === 'ArrowUp') {
              e.preventDefault();
              clampOptionIndex(optionsIndex - 1);
              return;
            }
            if (key === 'Enter' || key === 'Tab') {
              e.preventDefault();
              showOptions = false;
              selectedOption = options[optionsIndex] ?? '';
              return;
            }
          }}
          class="relative z-10 size-full bg-clip-padding px-2 text-left text-ctp-lavender"
          value={selectedOption}
        />
        <span class="absolute right-12 icon-[ri--arrow-down-s-line] size-5 h-full"></span>
        <!-- svelte-ignore a11y_consider_explicit_label -->
        <button
          onmousedown={async () => {
            if (selectedOption !== '') {
              response = submitOption(selectedOption);
            } else {
              response = Promise.resolve({ error: true, message: 'no option selected' });
            }
          }}
          class="flex h-full w-12 cursor-pointer items-center justify-center bg-jfgray-800 peer-focus:bg-jfgray-800"
        >
          {#await response}
            <span in:fade class="icon-[ri--loader-3-line] animate-spin text-ctp-lavender"></span>
          {:then response}
            <span
              in:fade
              class="icon-[ri--send-plane-line] {response.error === true
                ? 'text-ctp-red'
                : 'text-ctp-lavender'}"
            ></span>
          {/await}
        </button>
      </div>
    </label>
    {#if showOptions}
      <div
        transition:slide
        class="absolute top-full z-40 flex w-full flex-col border-2 border-ctp-lavender-50/50 bg-jfgray-900 p-2"
      >
        {#each options as option, i}
          <InputOption
            value={option}
            onSelect={(value: string) => {
              selectedOption = value;
              showOptions = false;
            }}
            index={i}
            selectedIndex={optionsIndex}
          />
        {/each}
      </div>
    {/if}
  </div>
  {#await response then response}
    <Response {response} />
  {/await}
</div>
