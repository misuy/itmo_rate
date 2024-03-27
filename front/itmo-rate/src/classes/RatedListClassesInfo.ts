export default class RatedListItemInfo {
    rating: number;
    name: string;
    tags: string[];

    constructor (rating: number, name: string, tags: string[]) {
        this.rating = rating;
        this.name = name;
        this.tags = tags;
    }
}