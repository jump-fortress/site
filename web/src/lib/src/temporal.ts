import { Temporal } from 'temporal-polyfill';

const units = new Map<'year' | 'month' | 'day' | 'hour' | 'minute', number>([
  ['year', 24 * 60 * 60 * 1000 * 365],
  ['month', (24 * 60 * 60 * 1000 * 365) / 12],
  ['day', 24 * 60 * 60 * 1000],
  ['hour', 60 * 60 * 1000],
  ['minute', 60 * 1000]
]);

const dateFormatter = new Intl.DateTimeFormat(undefined, {
  day: 'numeric',
  month: 'short',
  year: 'numeric'
});

const timeFormatter = new Intl.DateTimeFormat(undefined, {
  hour12: false,
  hour: '2-digit',
  minute: '2-digit',
  timeZoneName: 'shortOffset',
  timeZone: Temporal.Now.timeZoneId()
});

const relativeFormatter = new Intl.RelativeTimeFormat(undefined, { numeric: 'always' });

export function formatDate(time: number) {
  return dateFormatter.format(time).replace(',', '');
}

export function formatTime(time: number) {
  return timeFormatter.format(time).replace('GMT', 'UTC');
}

export function formatRelative(time: number) {
  const now = Temporal.Now.instant().epochMilliseconds;
  const diff = time - now;

  for (const [unit, ms] of units)
    if (Math.abs(diff) > ms) {
      return relativeFormatter.format(Math.round(diff / ms), unit);
    }
  return relativeFormatter.format(Math.round(diff / (60 * 1000)), 'minute');
}

export function DateToMs(date: string) {
  return Temporal.Instant.from(date).epochMilliseconds;
}
