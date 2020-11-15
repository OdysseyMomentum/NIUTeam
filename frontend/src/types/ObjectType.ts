export interface IObjectType {
  uoi: string;
  displayName: string;
  description?: string;
  geoCoordinates: {
    latitude: number;
    longitude: number;
  };
  streetName: string;
  streetNumber: string;
  zipcode: string;
  city: string;
  country: string;
}
