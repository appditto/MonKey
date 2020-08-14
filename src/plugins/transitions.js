import { cubicOut } from 'svelte/easing';

export const fadeAndScaleIn = (node, { delay = 0, duration = 500 }) => {
  return {
    delay,
    duration,
    css: (t) => {
      const eased = cubicOut(t);
      return `opacity: ${(eased * 1 / 2) + 0.5}; transform: scale(${(eased * 1 / 2) + 0.5})`
    },
  };
};

export const fadeAndScaleOut = (node, { delay = 0, duration = 200 }) => {
  return {
    delay,
    duration,
    css: (t) => {
      return `opacity: ${t}; transform: scale(${t * 1 / 2 + 0.5})`
    },
  };
};