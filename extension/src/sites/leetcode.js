const LeetCode = {
    name: "leetcode",

    matches() {
        return (
            window.location.hostname.includes("leetcode.com") &&
            this.extractSlug() !== null
        );
    },

    extractSlug() {
        const path = window.location.pathname;
        const parts = path.split("/").filter(Boolean);

        if (parts.length >= 2 && parts[0] === "problems") {
            return parts[1];
        }

        return null;
    },

    extractTitleFromSlug(slug) {
        if (!slug) return null;

        return slug
            .split("-")
            .map(word => word.charAt(0).toUpperCase() + word.slice(1))
            .join(" ");
    },

    async waitForDescription(timeout = 10000) {
        return new Promise((resolve, reject) => {
            const start = Date.now();

            const check = () => {
                const container = document.querySelector(
                    '[data-track-load="description_content"]'
                );

                if (container && container.querySelector("pre")) {
                    return resolve(container);
                }

                if (Date.now() - start > timeout) {
                    return reject("Timeout waiting for description");
                }

                requestAnimationFrame(check);
            };

            check();
        });
    },

    // testcases
    extractTestsFromDom() {
        const container = document.querySelector(
            '[data-track-load="description_content"]'
        );

        if (!container) return [];

        const pres = container.querySelectorAll("pre");
        const tests = [];

        pres.forEach(pre => {
            let input = null;
            let output = null;

            const nodes = [...pre.childNodes];

            for (let i = 0; i < nodes.length; i++) {
                const node = nodes[i];

                if (node.nodeName === "STRONG") {
                    const label = node.innerText.trim();

                    if (label === "Input:") {
                        input = this.collectTextUntilNextStrong(nodes, i + 1);
                    }

                    if (label === "Output:") {
                        output = this.collectTextUntilNextStrong(nodes, i + 1);
                    }
                }
            }

            if (input && output) {
                tests.push({ input, output });
            }
        });

        return tests;
    },


    collectTextUntilNextStrong(nodes, startIndex) {
        let text = "";

        for (let i = startIndex; i < nodes.length; i++) {
            const current = nodes[i];

            if (current.nodeName === "STRONG") {
                break;
            }

            text += current.textContent;
        }

        return text.trim();
    },

// language
extractLanguageFromDom() {
    const buttons = document.querySelectorAll('button[aria-haspopup="dialog"]');

    for (const btn of buttons) {
        const text = btn.innerText.trim();

        if (text && text.length <= 15 && this.isLikelyLanguage(text)) {
            return {
                display: text,
                normalized: this.normalizeLanguage(text)
            };
        }
    }

    return null;
},

isLikelyLanguage(text) {
    const known = [
        "C++",
        "Java",
        "Python",
        "Python3",
        "Go",
        "JavaScript",
        "TypeScript",
        "C",
        "C#",
        "Rust",
        "Kotlin"
    ];

    return known.includes(text);
},

normalizeLanguage(lang) {
    const map = {
        "C++": "cpp",
        "Python": "py",
        "Python3": "py",
        "Java": "java",
        "Go": "go",
        "JavaScript": "js",
        "TypeScript": "ts",
        "C#": "cs",
        "C": "c",
        "Rust": "rs",
        "Kotlin": "kt"
    };

    return map[lang] || lang.toLowerCase();
},


    async extractProblem() {
        const slug = this.extractSlug();
        if (!slug) return null;

        try {
            await this.waitForDescription();
        } catch (err) {
            console.warn("Description did not load in time:", err);
            return null;
        }

        const title = this.extractTitleFromSlug(slug);
        const tests = this.extractTestsFromDom();
        const languageData = this.extractLanguageFromDom();

        return {
            slug,
            title,
            url: window.location.href,
            source: this.name,
            language: languageData?.normalized || null,
            languageDisplay: languageData?.display || null,
            tests
        };
    }
};

export default LeetCode;
