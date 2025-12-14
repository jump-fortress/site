<script lang="ts">
  import SelectOption from '$lib/components/input/InputSelectOption.svelte';
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

  // response is a promise so #await can be used to handle pending state
  let response: Promise<InputError> = $derived(
    Promise.resolve({
      error: false,
      message: responseMessage
    })
  );
</script>

<!-- todo: judge slide transition, send request -->
<div class="relative flex h-12 w-80 items-center gap-2">
  <label
    for={label}
    class="relative mt-2 w-full border-2 border-jfgray-700 text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 hover:border-ctp-lavender-50/50 {selectedOption
      ? 'bg-jfgray-900'
      : 'bg-jfgray-800'}"
  >
    <span class="absolute -top-1 left-2 bg-jfgray-800 px-1 text-base leading-1">{label}</span>
    <div class="flex h-10">
      <input
        type="button"
        id={label}
        onfocus={() => {
          showOptions = true;
        }}
        onfocusout={() => {
          showOptions = false;
        }}
        class="size-full bg-clip-padding px-2 text-left"
        value={selectedOption}
      />
      <span class="mr-2 icon-[ri--arrow-down-s-line] size-6 h-full"></span>
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <button
        onclick={async () => {
          response = submitOption(selectedOption);
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
      {#each options as option}
        <SelectOption
          value={option}
          onSelect={(value: string) => {
            selectedOption = value;
            showOptions = false;
          }}
        />
      {/each}
    </div>
  {/if}
  {#await response then response}
    <Response {response} />
  {/await}
</div>
