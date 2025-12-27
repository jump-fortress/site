<script lang="ts">
  import InputOption from '$lib/components/input/InputOption.svelte';
  import Response from '$lib/components/input/Response.svelte';
  import { fade, slide } from 'svelte/transition';
  import Input from './Input.svelte';

  type Props = {
    label: string;
    options: string[];
    placeholder?: string;
    responseMessage?: string;
    submitOption: (val: string) => Promise<InputError>;
  };
  let { label, options, responseMessage = '', placeholder = '', submitOption }: Props = $props();

  let inputValue: string = $state(placeholder);
  let autofillIndex: number = $state(0);
  let autofillOptions: string[] = $state([]);
  let showOptions: boolean = $state(false);

  // response is a promise so #await can be used to handle pending state
  let response: Promise<InputError> = $state(
    Promise.resolve({
      error: false,
      message: responseMessage
    })
  );

  function clampAutofillIndex(n: number) {
    autofillIndex = n < 0 ? 0 : n > autofillOptions.length - 1 ? autofillOptions.length - 1 : n;
  }

  $effect(() => {
    autofillOptions = options.filter((option) => option.includes(inputValue));
  });
</script>

<div class="flex items-center gap-2">
  <div class="relative flex h-12 w-80 items-center">
    <label
      for={label}
      class="relative mt-1 w-full border-2 border-jfgray-700 text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 focus-within:bg-jfgray-900 hover:border-ctp-lavender-50/50
      {showOptions ? 'rounded-t-md' : 'rounded-md'} 
      {inputValue ? 'bg-jfgray-900' : 'bg-jfgray-800'}">
      <span class="absolute -top-1 left-2 bg-jfgray-800 px-1 text-sm leading-1">{label}</span>
      <div class="relative flex h-10">
        <input
          type="text"
          id={label}
          onmousedown={() => {
            showOptions = true;
          }}
          onfocus={() => {
            showOptions = true;
          }}
          onfocusout={() => {
            showOptions = false;
          }}
          onkeydown={(e) => {
            const key = e.key;
            if (key === 'ArrowDown') {
              e.preventDefault();
              clampAutofillIndex(autofillIndex + 1);
              return;
            }
            if (key === 'ArrowUp') {
              e.preventDefault();
              clampAutofillIndex(autofillIndex - 1);
              return;
            }
            if (key === 'Enter' || key === 'Tab') {
              e.preventDefault();
              showOptions = false;
              inputValue = autofillOptions[autofillIndex] ?? '';
              return;
            }
            showOptions = true;
          }}
          bind:value={inputValue}
          class="relative z-10 size-full px-2 text-left text-ctp-lavender" />
        <span class="absolute right-12 icon-[ri--arrow-down-s-line] size-5 h-full cursor-text"
        ></span>
        <!-- svelte-ignore a11y_consider_explicit_label -->
        <button
          onmousedown={async () => {
            response = submitOption(inputValue);
          }}
          class="flex h-full w-10 cursor-pointer items-center justify-center">
          {#await response}
            <span in:fade class="icon-[ri--loader-3-line] animate-spin text-ctp-lavender"></span>
          {:then response}
            <span
              in:fade
              class="icon-[ri--send-plane-line] {response.error === true
                ? 'text-ctp-red'
                : 'text-ctp-lavender'}"></span>
          {/await}
        </button>
      </div>
    </label>
    {#if showOptions}
      <div
        class="absolute top-full z-40 flex max-h-60 w-7/8 flex-col overflow-x-hidden overflow-y-auto rounded-b-md border-2 border-t-0 border-ctp-lavender-50/50 bg-jfgray-900">
        {#if autofillOptions.length > 0}
          {#each autofillOptions as option, i}
            <InputOption
              value={option}
              onSelect={(value: string) => {
                inputValue = value;
                showOptions = false;
              }}
              index={i}
              selectedIndex={autofillIndex} />
          {/each}
        {/if}
      </div>
    {/if}
  </div>
  {#await response then response}
    <Response {response} />
  {/await}
</div>
