// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Create as $Create} from "@wailsio/runtime";

export class Author {
    "Id": number;
    "Name": string;
    "Posts": Post[];

    /** Creates a new Author instance. */
    constructor($$source: Partial<Author> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("Name" in $$source)) {
            this["Name"] = "";
        }
        if (!("Posts" in $$source)) {
            this["Posts"] = [];
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Author instance from a string or object.
     */
    static createFrom($$source: any = {}): Author {
        const $$createField2_0 = $$createType1;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("Posts" in $$parsedSource) {
            $$parsedSource["Posts"] = $$createField2_0($$parsedSource["Posts"]);
        }
        return new Author($$parsedSource as Partial<Author>);
    }
}

export class Post {
    "Id": number;
    "Text": string;
    "Deadline": number;
    "CreatedAt": number;
    "AuthorId": number;
    "Author": Author;

    /** Creates a new Post instance. */
    constructor($$source: Partial<Post> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("Text" in $$source)) {
            this["Text"] = "";
        }
        if (!("Deadline" in $$source)) {
            this["Deadline"] = 0;
        }
        if (!("CreatedAt" in $$source)) {
            this["CreatedAt"] = 0;
        }
        if (!("AuthorId" in $$source)) {
            this["AuthorId"] = 0;
        }
        if (!("Author" in $$source)) {
            this["Author"] = (new Author());
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Post instance from a string or object.
     */
    static createFrom($$source: any = {}): Post {
        const $$createField5_0 = $$createType2;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("Author" in $$parsedSource) {
            $$parsedSource["Author"] = $$createField5_0($$parsedSource["Author"]);
        }
        return new Post($$parsedSource as Partial<Post>);
    }
}

// Private type creation functions
const $$createType0 = Post.createFrom;
const $$createType1 = $Create.Array($$createType0);
const $$createType2 = Author.createFrom;
