export function validateSite(site) {
    if (typeof site.name !== "string") throw Error("Invalid site: name");
    if (typeof site.matches !== "function") throw Error("Invalid site: matches()");
    if (typeof site.extractProblem !== "function") throw Error("Invalid site: extractProblem()");
}
