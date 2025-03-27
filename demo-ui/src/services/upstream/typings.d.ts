declare namespace API {
    type Upstream = {
        name: string;
        type: string;
        nodes: {
            host: string;
            port: number;
            weight: number;
        }[];
        describe: string;
    }
    type UpstreamList = {
        data?: Upstream[];
        total?: number;
        success?: boolean;
    }
}