
declare namespace API {

  type Service = {
    name: string;
    namespace?: string;
    labels?: { [key: string]: string};
    annotations?: { [key: string]: string};
    selector?: { [key: string]: string};
    ports: Array<{
        name?: string;
        protocol: 'TCP' | 'UDP';
        port: number;
        targetPort?: number | string;
    }>;
    type: 'ClusterIP' | 'NodePort' | 'LoadBalancer' | 'ExternalName';
    clusterIP?: string;
    externalIPs?: string[];
    externalName?: string;
  }
   
  type ServiceItem = {
    name: string;
    namespace?: string;
    clusterIP?: string;
    externalIPs?: string[];
    ports?: string;
  }
  type ServiceList = {
    data?: ServiceItem[];
    total?: number;
    success?: boolean;
  }
  type Container = {
    name: string;
    image: string;
    ports?: Array<{
      containerPort: number;
      protocol?: 'TCP' | 'UDP';
    }>;
    resources?: {
      limits?: { [key: string]: string };
      requests?: { [key: string]: string };
    };
  }
  type Deployment = {
      name: string;
      namespace?: string;
      labels?: { [key: string]: string};
      annotations?: { [key: string]: string};
      replicas?: number;
      selector?: { [key: string]: string};
      strategy?: {
        rollingUpdate?: {
          maxSurge?: number | string;
          maxUnavailable?: number | string;
        },
        type?: 'RollingUpdate' | 'Recreate',
      };
      containers: Container[];
      
  }
  type DeploymentList = {
      data?: Deployment[];
      total?: number;
      success?: boolean;
  }
}