import LeetCode from "./leetcode.js";
import {validateSite} from "./siteValidator.js"

export const registeredSites  = [
    LeetCode
];


registeredSites.forEach(validateSite);

export const sites = registeredSites