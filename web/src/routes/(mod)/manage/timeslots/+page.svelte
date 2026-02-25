<script lang="ts">
  import { Client } from '$lib/api/api';
  import Table from '$lib/components/display/table/Table.svelte';
  import TemporalDate from '$lib/components/display/TemporalDate.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Errors from '$lib/components/input/Errors.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Select from '$lib/components/input/Select.svelte';
  import Content from '$lib/components/layout/Content.svelte';
  import { datetimeToMs, formatTime, validDateTime } from '$lib/helpers/temporal';
  import { ApiPaths, type MotwTimeslot } from '$lib/schema';
  import { Temporal } from 'temporal-polyfill';

  let oerror: OpenAPIError = $state(undefined);

  let refreshTimeslots: boolean = $state(false);

  let timeslot_date: string = $state(Temporal.Now.plainDateISO().toString());
  $inspect(oerror);
  let timeInput: string = $state('');
  const timeslotDatetime: string = $derived(validDateTime(`${timeslot_date}T${timeInput}:00Z`));
  let id: number = $state(1);

  // function loadTimeslot(t: MotwTimeslot) {
  //   id = t.id;
  //   timeInput = t.starts_at.substring(t.starts_at.indexOf('T') + 1, t.starts_at.indexOf('Z') - 3);
  //   timeslot_date = t.starts_at.substring(0, t.starts_at.indexOf('T'));
  // }
</script>

<Content>
  <Errors {oerror} />
  {#key refreshTimeslots}
    {#await Client.GET(ApiPaths.get_timeslot)}
      <span></span>
    {:then { data: timeslotInfo }}
      {#if timeslotInfo?.timeslots}
        <div class="flex w-full">
          <div class="flex grow flex-col">
            <span class="text-primary"
              >selected: {formatTime(
                datetimeToMs(timeslotInfo.timeslots.at(id - 1)!.starts_at)
              )}</span>
            <Select
              type="button"
              label="select timeslot"
              options={timeslotInfo.timeslots.map((t) => formatTime(datetimeToMs(t.starts_at)))}
              onsubmit={async (_, i) => {
                id = i + 1;
                return true;
              }} />
            <Input
              label="new start time"
              type="time"
              bind:value={timeInput}
              withSubmit={false}
              onsubmit={async (value) => {
                return true;
                // let resp = await Client.POST(ApiPaths.update_timeslot, {body: })
                //       oerror = resp.error;
                //       if (resp.response.ok) {
                //         refreshPR = !refreshPR;
                //       }
                //       return resp.response.ok;}}
                //
              }} />
          </div>
          <div class="flex flex-col">
            <span
              >input timezone <span class="text-primary">UTC</span>
              <span></span>
            </span>
            <span
              >your timezone <span class="text-primary"
                >{Temporal.Now.timeZoneId()} ({Temporal.Now.zonedDateTimeISO().offset} UTC)</span
              ></span>
          </div>
        </div>
      {/if}

      <Button
        onsubmit={async () => {
          const resp = await Client.POST(ApiPaths.update_timeslot, {
            body: {
              id: id,
              starts_at: timeslotDatetime
            }
          });
          oerror = resp.error;
          if (resp.response.ok) {
            refreshTimeslots = !refreshTimeslots;
          }
          return resp.response.ok;
        }}><span>submit</span></Button>
    {/await}
  {/key}
</Content>
