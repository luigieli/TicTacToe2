import styles from "./LayoutBorder.module.css";

export default function LayoutBorder() {
  return (
    <>
      <div className={`${styles.layoutBorder} ${styles.borderBottom}`} />
      <div className={`${styles.layoutBorder} ${styles.borderTop}`} />
    </>
  );
}
