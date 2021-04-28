---
title: "Angular Core Utils"
date: 2021-01-20T19:56:17+08:00
draft: false
---

```ts
/**
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE file at https://github.com/NG-ZORRO/ng-zorro-antd/blob/master/LICENSE
 */

export function toArray<T>(value: T | T[]): T[] {
  let ret: T[];
  if (value == null) {
    ret = [];
  } else if (!Array.isArray(value)) {
    ret = [value];
  } else {
    ret = value;
  }
  return ret;
}

export function arraysEqual<T>(array1: T[], array2: T[]): boolean {
  if (!array1 || !array2 || array1.length !== array2.length) {
    return false;
  }

  const len = array1.length;
  for (let i = 0; i < len; i++) {
    if (array1[i] !== array2[i]) {
      return false;
    }
  }
  return true;
}

export function shallowCopyArray<T>(source: T[]): T[] {
  return source.slice();
}



export function getPercent(min: number, max: number, value: number): number {
  return ((value - min) / (max - min)) * 100;
}

export function getPrecision(num: number): number {
  const numStr = num.toString();
  const dotIndex = numStr.indexOf('.');
  return dotIndex >= 0 ? numStr.length - dotIndex - 1 : 0;
}

export function ensureNumberInRange(num: number, min: number, max: number): number {
  if (isNaN(num) || num < min) {
    return min;
  } else if (num > max) {
    return max;
  } else {
    return num;
  }
}

export function isNumberFinite(value: NzSafeAny): boolean {
  return typeof value === 'number' && isFinite(value);
}

export function toDecimal(value: number, decimal: number): number {
  return Math.round(value * Math.pow(10, decimal)) / Math.pow(10, decimal);
}

export function sum(input: number[], initial: number = 0): number {
  return input.reduce((previous: number, current: number) => previous + current, initial);
}



/**
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE file at https://github.com/NG-ZORRO/ng-zorro-antd/blob/master/LICENSE
 */

/**
 * Much like lodash.
 */
export function padStart(toPad: string, length: number, element: string): string {
  if (toPad.length > length) {
    return toPad;
  }

  const joined = `${getRepeatedElement(length, element)}${toPad}`;
  return joined.slice(joined.length - length, joined.length);
}

export function padEnd(toPad: string, length: number, element: string): string {
  const joined = `${toPad}${getRepeatedElement(length, element)}`;
  return joined.slice(0, length);
}

export function getRepeatedElement(length: number, element: string): string {
  return Array(length).fill(element).join('');
}