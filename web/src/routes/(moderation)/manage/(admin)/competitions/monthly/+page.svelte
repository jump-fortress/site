<script lang="ts">
  import DataSection from '$lib/components/DataSection.svelte';
  import InputSelect from '$lib/components/input/InputSelect.svelte';
  import { Temporal } from 'temporal-polyfill';
  import { divisions, compareDivisions } from '$lib/divisions';
  import InputAutofill from '$lib/components/input/InputAutofill.svelte';
  import CompetitionManagePreview from '$lib/components/CompetitionManagePreview.svelte';
  import type { Map } from '$lib/schema.js';
  import Input from '$lib/components/input/Input.svelte';

  let validTime = $state(true);

  let { data } = $props();

  let maps: Map[] = $state([]);
  let mapNames: string[] = $derived(maps.map((m) => m.name));

  let competitionMaps: MapWithDivisions[] = $state([]);
  let competitionMapNames: string[] = $derived(competitionMaps.map((m) => m.name));
  let competitionDivisions: string[] = $state([]);

  let prizepool: number = $state(0);

  function addCompetitionMap(map: Map | undefined) {
    if (!map || !mapNames.includes(map.name)) {
      return { error: true, message: 'invalid map name' };
    }
    if (competitionMapNames.includes(map.name)) {
      return { error: true, message: 'map already added' };
    }
    competitionMaps.push({ name: map.name, divisions: [] });
    return { error: false, message: '' };
  }

  maps = (await data.maps) ?? [];

  // todo: cleanup / componentize
</script>

{#await data.maps then maps}
  {#if maps}
    <DataSection title={'Create'}>
      <!-- date & time -->
      <div class="flex h-12 items-center gap-2">
        <label
          for="date"
          class="relative mt-2 w-50 rounded-md border-2 border-jfgray-700 text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 focus-within:bg-jfgray-900 hover:border-ctp-lavender-50/50">
          <span class="absolute -top-1 left-2 bg-jfgray-800 px-1 text-sm leading-1">date</span>
          <div class="flex h-10">
            <input
              class="size-full px-2 text-ctp-lavender"
              type="date"
              id="date"
              value={Temporal.Now.plainDateISO().toString()} />
          </div>
        </label>
        <label
          for="time"
          class="relative mt-2 w-50 rounded-md border-2 border-jfgray-700 text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 focus-within:bg-jfgray-900 hover:border-ctp-lavender-50/50 has-invalid:border-ctp-red">
          <span class="absolute -top-1 left-2 bg-jfgray-800 px-1 text-sm leading-1">time</span>
          <div class="flex h-10">
            <input
              oninput={(e) => {
                const valid = e.currentTarget.checkValidity();
                validTime = valid;
              }}
              class="size-full px-2 text-ctp-lavender"
              type="time"
              id="time"
              step="1800"
              value={Temporal.Now.plainTimeISO()
                .round('hour')
                .toString({ smallestUnit: 'minute' })} />
          </div>
        </label>
        {#if !validTime}
          <span>please set a time in steps of 30 minutes (:00 or :30)</span>
        {/if}
        <div class="ml-auto">
          <span class="text-ctp-lavender">timezone: </span>
          <span>{Temporal.Now.timeZoneId()}</span>
        </div>
      </div>
      <!-- add divisions & maps -->
      <Input
        label={'prizepool'}
        submitInput={async (val: string) => {
          const keys = parseInt(val);
          if (keys && keys > 0) {
            prizepool = keys;
          }
          return { error: false, message: '' };
        }} />

      <InputSelect
        label={'class'}
        options={['Soldier', 'Demo']}
        submitOption={async (val: string) => {
          return { error: false, message: '' };
        }} />

      <InputAutofill
        label={'maps'}
        options={maps.map((m) => m.name)}
        submitOption={async (val: string) => {
          return addCompetitionMap(maps.find((m) => m.name == val));
        }} />

      <!-- select divisions for maps -->
      <!-- {#each competitionMaps as map}
        <div class="flex items-center gap-2">
          <span>{map.name}</span>
          <InputSelect
            label={'select divisions'}
            options={divisions.concat('None')}
            submitOption={async (val: string) => {
              // if none, remove all divs
              if (val === 'None') {
                competitionDivisions.filter((div) => map.divisions.includes(div));
                map.divisions = [];
              } else {
                if (competitionDivisions.includes(val)) {
                  // remove this division from the competitionMap.divisions that has it
                  // todo: shouldn't i just use a hashmap.... it works though
                  const cmi = competitionMaps.findIndex(({ divisions }) => divisions.includes(val));
                  const cmdi = competitionMaps.at(cmi)?.divisions.findIndex((div) => div === val);
                  if (cmdi !== undefined) {
                    competitionMaps.at(cmi)?.divisions.splice(cmdi, 1);
                  }
                }
                map.divisions.push(val);
                competitionDivisions.push(val);
                map.divisions.sort(compareDivisions);
              }
              return { error: false, message: '' };
            }}
          />
        </div>
      {/each} -->

      {#if competitionMaps.length}
        <div class="grid grid-cols-3 gap-2">
          {#each divisions as division}
            <InputSelect
              label={division.toLowerCase() + ' map'}
              options={competitionMapNames}
              submitOption={async (val: string) => {
                const map = competitionMaps.find((m) => m.name === val);
                if (!map) return { error: false, message: '' };
                // if none, remove all divs
                if (competitionDivisions.includes(division)) {
                  // remove this division from the competitionMap.divisions that has it
                  // todo: shouldn't this use a hashmap or something or other..?
                  // note: originally written to add a division to a map, then frankensteined to add maps to divisions
                  const cmi = competitionMaps.findIndex(({ divisions }) =>
                    divisions.includes(division)
                  );
                  const cmdi = competitionMaps
                    .at(cmi)
                    ?.divisions.findIndex((div) => div === division);
                  if (cmdi !== undefined) {
                    competitionMaps.at(cmi)?.divisions.splice(cmdi, 1);
                  }
                }
                map.divisions.push(division);
                competitionDivisions.push(division);
                map.divisions.sort(compareDivisions);
                return { error: false, message: '' };
              }} />
          {/each}
        </div>
      {/if}

      <CompetitionManagePreview maps={competitionMaps} />
    </DataSection>
  {/if}
{/await}
