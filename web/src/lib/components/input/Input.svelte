<script lang="ts">
  import { fade } from 'svelte/transition';
  import Response from '$lib/components/input/Response.svelte';

  type Props = {
    label: string;
    placeholder?: string;
    responseMessage?: string;
    submitInput: (val: string) => Promise<InputError>;
  };

  let { label, placeholder = '', responseMessage = '', submitInput }: Props = $props();
  let value: string = $state('');

  // response is a promise so #await can be used to handle pending state
  let response: Promise<InputError> = $state(
    Promise.resolve({
      error: false,
      message: responseMessage
    })
  );
</script>

<!-- input & response container -->
<div class="flex items-center gap-2">
  <label
    for={label}
    class="relative mt-1 w-80 rounded-md border-2 border-jfgray-700 bg-jfgray-800 bg-clip-content text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 focus-within:bg-jfgray-900 hover:border-ctp-lavender-50/50">
    <!-- floating label -->
    <span class="absolute -top-1 left-2 bg-jfgray-800 px-1 text-sm leading-1">{label}</span>
    <!-- input container -->
    <div class="flex h-10">
      <input
        bind:value
        {placeholder}
        onkeypress={async (event: KeyboardEvent) => {
          if (event.key === 'Enter' && value) {
            response = submitInput(value);
          }
        }}
        class=" peer size-full px-2 text-ctp-lavender transition-colors focus:outline-0"
        id={label}
        type="text" />
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <button
        onmousedown={async () => {
          response = submitInput(value);
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
  {#await response then response}
    <Response {response} />
  {/await}
</div>
